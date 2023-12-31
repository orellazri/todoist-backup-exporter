<div align="center">

# Todoist Backup Exporter

This project allows you to export your last Todoist backup to a zip file.

[![Build and Push Docker Image](https://github.com/orellazri/todoist-backup-exporter/actions/workflows/release.yml/badge.svg)](https://github.com/orellazri/todoist-backup-exporter/actions/workflows/release.yml)

</div>

## Running

### Docker (Recommended)

```bash
docker run -- rm \
    -e TODOIST_TOKEN=... \
    -e ... \
    -v $(pwd):/exports \
    reaperberri/todoist-backup-exporter
```

Note: You need to mount a volume to the `/exports` directory to get the file out of the container (the example comand above mounts the current directory).

### Go

```bash
# export environment variables here like so:
# export DEST_DIR=...
go run .
```

## Environment variables

| Required | Name                       | Description                                                                                    | Default value |
| -------- | -------------------------- | ---------------------------------------------------------------------------------------------- | ------------- |
| Yes      | TODOIST_TOKEN              | Your Todoist API token                                                                         |               |
| No       | DEST_DIR                   | Destination directory to download the file to (Don't need to change if running in a container) | /exports      |
| No       | DEST_FILENAME_INCLUDE_DATE | Whether to include the backup date in the file name                                            | false         |

To get the Todoist API token, go to [Todoist Integrations](https://todoist.com/prefs/integrations), click the "Developer" tab and copy the token from the API token section.
