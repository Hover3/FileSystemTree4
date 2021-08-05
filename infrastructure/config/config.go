package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	RootDirectory          string
	EnabledExtensions      []string

	EnableColorText        bool
	FoldersColor           string
	FilesColor             string
	FileStatsColor         string
	TreeBranchColor        string
	TreeExpandColor        string
	TreeCollapseColor      string
	WarningColor           string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		RootDirectory:          getEnv("ROOT_DIRECTORY", ""),
		EnabledExtensions:      getEnvAsSlice("ENABLED_EXTENSIONS", []string{".go", ".md"}, ","),

		EnableColorText:        getEnvAsBool("ENABLE_COLOR_TEXT", false),
		FoldersColor:           getEnv("FOLDERS_COLOR", "reset"),
		FilesColor:             getEnv("FILES_COLOR", "reset"),
		FileStatsColor:         getEnv("FILE_STATS_COLOR", "reset"),
		TreeBranchColor:        getEnv("TREE_BRANCH_COLOR", "reset"),
		TreeExpandColor:        getEnv("TREE_EXPAND_COLOR", "reset"),
		TreeCollapseColor:      getEnv("TREE_COLLAPSE_COLOR", "reset"),
		WarningColor:           getEnv("WARNING_COLOR", "reset"),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
