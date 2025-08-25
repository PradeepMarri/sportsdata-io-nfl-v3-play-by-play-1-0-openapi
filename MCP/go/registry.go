package main

import (
	"github.com/nfl-v3-play-by-play/mcp-server/config"
	"github.com/nfl-v3-play-by-play/mcp-server/models"
	tools_format "github.com/nfl-v3-play-by-play/mcp-server/tools/format"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_format.CreatePlaybyplayTool(cfg),
		tools_format.CreatePlaybyplaydeltaTool(cfg),
		tools_format.CreatePlaybyplaysimulationTool(cfg),
	}
}
