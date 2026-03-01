# Что это?

Линтер логов для Go, совместимый с [golangci-lint](https://golangci-lint.run/). Собирается как модуль для golang-cli по "[автоматическому методу](https://golangci-lint.run/docs/plugins/module-plugins/#the-automatic-way)".

Собрать можно командой (соберется кастомный бинарный файл ./custom-gcl):
```
golangci-lint custom
```

Для использования в проекте нужно использовать файл [.golangci.yml](.golangci.yml) в корне проекта. 

# Тесты

Можно запустить с помощью:
```
go test
```

# CI
Настроен на push и pull_request. Запускает тесты и тестовую сборку плагина.

# Конфигурация

Используется стандартная конфигурация .golangci.yml

Линтер позволяет задавать настройках кастомные паттерны для проверки на чувствительные данные. Паттерны нужно задавать в нижнем регистре.

# Пример использования на моем репозитории

<img width="848" height="119" alt="image" src="https://github.com/user-attachments/assets/07c8520d-79cb-4abf-bd7a-9af2330b608d" />

