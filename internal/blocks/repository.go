package blocks

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type BlocksRepository interface {
	GetBlocks() ([]string, error)
	AddBlock(link string) error
	RemoveBlock(link string) error
}

type fileBlocksRepository struct {
	filePath string
}

func BlocksFilePath() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		dir = os.TempDir()
	}
	return filepath.Join(dir, "pomodoro_blocks.json")
}

func NewBlocksRepository(filePath string) BlocksRepository {
	return &fileBlocksRepository{filePath: filePath}
}

func (r *fileBlocksRepository) GetBlocks() ([]string, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var blocks []string
	if err := json.Unmarshal(data, &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

func (r *fileBlocksRepository) AddBlock(link string) error {
	blocks, err := r.GetBlocks()
	if err != nil {
		return err
	}

	// Check if already exists
	for _, b := range blocks {
		if b == link {
			return nil // already blocked
		}
	}

	blocks = append(blocks, link)
	return r.save(blocks)
}

func (r *fileBlocksRepository) RemoveBlock(link string) error {
	blocks, err := r.GetBlocks()
	if err != nil {
		return err
	}

	var newBlocks []string
	for _, b := range blocks {
		if b != link {
			newBlocks = append(newBlocks, b)
		}
	}

	return r.save(newBlocks)
}

func (r *fileBlocksRepository) save(blocks []string) error {
	data, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}
