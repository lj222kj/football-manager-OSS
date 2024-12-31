# football-manager-OSS

## Team Service

The Team Service is responsible for managing all operations related to football teams. This includes creating, updating, and retrieving team data, as well as managing the relationship between teams and their players.

Create and manage team details.

Associate players with a team.

Fetch team information and roster.

## Player Service

The Player Service manages player attributes, positions, and career progression.

Create and update player profiles (e.g., attributes, age, position).

Retrieve player details (e.g., position, skill levels).

Manage player transfers and team assignments.

## Simulation Service

The Simulation Service is responsible for simulating football matches, tournaments, and other events. It calculates match results, tracks player performance, and generates game scenarios.

Simulate individual matches, including detailed events (e.g., goals, cards).

Generate season-long simulations, including league standings.

Provide real-time or step-by-step match updates for the frontend.

## Configuration
The application uses a centralized configuration to manage environment variables. The application automatically load relevant values from the environment and some sane defaults on start.

APP_NAME, ENV: General application info (app name, environment).
HOST, PORT: Network binding information for running the api.
REDIS_HOST, REDIS_PORT, REDIS_PASSWORD: Redis connection details, specifying host address, port, and password.

## Redis

Redis is used as the main data store.
