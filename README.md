# Go-Timezone

A Go application demonstrating proper timezone handling in web applications using GORM. This project showcases best practices for storing and retrieving datetime data while respecting user timezones.

## 🌍 Problem Solved

Common timezone-related challenges this project addresses:
- Storing dates in a standardized timezone (UTC) in the database
- Converting user input times to UTC before storage
- Retrieving and displaying data in user's local timezone
- Handling timezone conversions in queries and filters
- Maintaining data consistency across different timezones

## 🏗️ Project Structure

```
go-timezone/
├── api/           # API handlers and routes
├── bootstrap/     # Application bootstrapping
├── cmd/          # Command line tools
├── config/       # Configuration management
├── constants/    # Global constants
├── docker/       # Docker configurations
├── lib/          # Common libraries
├── migrations/   # Database migrations
├── models/       # Data models
├── .env.example
├── docker-compose.yml
├── Makefile
├── dbconfig.yml
└── main.go
```

## 🚀 Getting Started

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-timezone.git
cd go-timezone
```

2. Set up environment variables:
```bash
cp .env.example .env
```

3. Start the application using Docker:
```bash
docker-compose up -d
```

### Running Migrations
```bash
make migrate-up
```

## 💡 Key Features

1. **UTC Storage**
   - All datetime values are automatically converted to UTC before storage
   - Ensures consistent data storage regardless of server or user timezone

2. **Timezone-Aware Queries**
   - Date filtering respects user's timezone
   - Query results are automatically converted to user's timezone

3. **Transaction Support**
   - Proper transaction handling for data consistency
   - Atomic operations for timezone-sensitive data

## 🔍 Common Issues Solved

1. Date filtering across timezone boundaries
2. Inconsistent timestamp storage
3. Query result timezone mismatches
4. Transaction atomicity with timezone conversions

## 📄 License

This project is licensed under the MIT License.
