package loglint

import (
	"log/slog"

	"go.uber.org/zap"
)

func TestSlogLower() {
	slog.Info("Starting server on port 8080")   // want "log check failed: message should start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log check failed: message should start with a lowercase letter"

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
}

func TestSlogEnglish() {
	slog.Info("запуск сервера")                    // want "log check failed: message should be in english"
	slog.Error("ошибка подключения к базе данных") // want "log check failed: message should be in english"

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
}

func TestSlogSpecial() {
	slog.Info("server started! 🚀")                // want "log check failed: message should not contain special symbols"
	slog.Error("connection failed!!!")            // want "log check failed: message should not contain special symbols"
	slog.Warn("warning: something went wrong...") // want "log check failed: message should not contain special symbols"

	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
}

func TestSlogCritical() {
	password := "12345678"
	apiKey := "veryverysecretkeydonotshareoreverythingwillexplode"
	token := "veryverysecrettokendonotshareoreverythingwillexplode"

	slog.Info("user password ", "", password) // want "log check failed: message should not contain potential secrets"
	slog.Debug("api_key=" + apiKey)           // want "log check failed: message should not contain potential secrets"
	slog.Info("token: " + token)              // want "log check failed: message should not contain potential secrets"

	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token validated")
}

func TestZapLower() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	sugar.Info("Starting server on port 8080")   // want "log check failed: message should start with a lowercase letter"
	sugar.Error("Failed to connect to database") // want "log check failed: message should start with a lowercase letter"

	sugar.Info("starting server on port 8080")
	sugar.Error("failed to connect to database")
}

func TestZapEnglish() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	sugar.Info("запуск сервера")                    // want "log check failed: message should be in english"
	sugar.Error("ошибка подключения к базе данных") // want "log check failed: message should be in english"

	sugar.Info("server started")
	sugar.Error("connection failed")
	sugar.Warn("something went wrong")
}

func TestZapSpecial() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	sugar.Info("server started! 🚀")                // want "log check failed: message should not contain special symbols"
	sugar.Error("connection failed!!!")            // want "log check failed: message should not contain special symbols"
	sugar.Warn("warning: something went wrong...") // want "log check failed: message should not contain special symbols"

	sugar.Info("server started")
	sugar.Error("connection failed")
	sugar.Warn("something went wrong")
}

func TestZapCritical() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	password := "12345678"
	apiKey := "veryverysecretkeydonotshareoreverythingwillexplode"
	token := "veryverysecrettokendonotshareoreverythingwillexplode"

	sugar.Info("user password ", "", password) // want "log check failed: message should not contain potential secrets"
	sugar.Debug("api_key=" + apiKey)           // want "log check failed: message should not contain potential secrets"
	sugar.Info("token: " + token)              // want "log check failed: message should not contain potential secrets"

	sugar.Info("user authenticated successfully")
	sugar.Debug("api request completed")
	sugar.Info("token validated")
}
