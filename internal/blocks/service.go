package blocks

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	hostsMarkerBegin = "# POMODORO-CLI-BEGIN"
	hostsMarkerEnd   = "# POMODORO-CLI-END"
)

type BlocksService interface {
	AddBlock(link string) error
	RemoveBlock(link string) error
	GetBlocks() ([]string, error)
	ApplyBlocks() error
	RemoveAppliedBlocks() error
}

type blocksService struct {
	repo BlocksRepository
}

func NewBlocksService(repo BlocksRepository) BlocksService {
	return &blocksService{repo: repo}
}

func (s *blocksService) AddBlock(link string) error {
	// Simple validation
	url := link
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("invalid link: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not reach the link (is it correct?): %w", err)
	}
	defer resp.Body.Close()
	
	// As many sites return 403 or 404 due to anti-bots or non-existent routes,
	// any HTTP response means the domain exists. The real error
	// occurs in client.Do when DNS resolution fails.
	
	return s.repo.AddBlock(link)
}

func (s *blocksService) RemoveBlock(link string) error {
	return s.repo.RemoveBlock(link)
}

func (s *blocksService) GetBlocks() ([]string, error) {
	return s.repo.GetBlocks()
}

func hostsFilePath() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32\drivers\etc\hosts`
	}
	return "/etc/hosts"
}

func (s *blocksService) ApplyBlocks() error {
	blocks, err := s.repo.GetBlocks()
	if err != nil || len(blocks) == 0 {
		return nil // nothing to block
	}

	path := hostsFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read hosts file (requires admin/sudo?): %v", err)
	}

	content := string(data)
	if strings.Contains(content, hostsMarkerBegin) {
		err = s.RemoveAppliedBlocks()
		if err != nil {
			return err
		}
		data, _ = os.ReadFile(path)
		content = string(data)
	}

	var sb strings.Builder
	sb.WriteString("\n" + hostsMarkerBegin + "\n")
	for _, b := range blocks {
		domain := strings.TrimPrefix(b, "https://")
		domain = strings.TrimPrefix(domain, "http://")
		domain = strings.TrimSuffix(domain, "/")

		sb.WriteString(fmt.Sprintf("127.0.0.1 %s\n", domain))
		if !strings.HasPrefix(domain, "www.") {
			sb.WriteString(fmt.Sprintf("127.0.0.1 www.%s\n", domain))
		}
	}
	sb.WriteString(hostsMarkerEnd + "\n")

	content += sb.String()

	err = os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("could not write to hosts file (requires admin/sudo?): %v", err)
	}
	return nil
}

func (s *blocksService) RemoveAppliedBlocks() error {
	path := hostsFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return err 
	}

	content := string(data)
	if !strings.Contains(content, hostsMarkerBegin) {
		return nil 
	}

	lines := strings.Split(content, "\n")
	var newLines []string
	inBlock := false

	for _, line := range lines {
		trimLine := strings.TrimSpace(line)
		if trimLine == hostsMarkerBegin {
			inBlock = true
			continue
		}
		if trimLine == hostsMarkerEnd {
			inBlock = false
			continue
		}
		if !inBlock {
			newLines = append(newLines, line)
		}
	}

	newContent := strings.Join(newLines, "\n")
	newContent = strings.TrimRight(newContent, "\n") + "\n"

	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("could not write to hosts file to remove blocks (requires admin/sudo?): %v", err)
	}
	return nil
}
