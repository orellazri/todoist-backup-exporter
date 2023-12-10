# Todoist Backup Exporter

This project allows your to exports your last Todoist backup to a zip file.

## Usage

### Environment variables

| Required | Name                       | Description                                         | Default value |
| -------- | -------------------------- | --------------------------------------------------- | ------------- |
| Yes      | TODOIST_TOKEN              | Your Todoist API token                              |               |
| No       | DEST_DIR                   | Destination directory to download the file to       | .             |
| No       | DEST_FILENAME_INCLUDE_DATE | Whether to include the backup date in the file name | false         |

To get the Todoist API token, go to [Todoist Integrations](https://todoist.com/prefs/integrations), click the "Developer" tab and copy the token from the API token section.

### Running

## Development

1. Clone this repository
2. Run `go run .`
