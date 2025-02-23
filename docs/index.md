# Backstage Plugin

## Get involved

* Open source: [GitHub Repository](https://github.com/chussenot/steampipe-plugin-backstage)
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)

## Authentication

The Backstage plugin requires a host URL and an API token for authentication.

### Generating an API Token

To generate an API token:

1. Log in to your Backstage instance
2. Navigate to your user settings
3. Generate a new API token
4. Copy the token value (it will only be shown once)

For more information about Backstage authentication, see:

- [Backstage Authentication](https://backstage.io/docs/auth/)
- [Backstage Tokens](https://backstage.io/docs/auth/tokens)

### Required Permissions

The API token needs the following permissions to query the catalog:

- `catalog.entity.read`
- `catalog.location.read`

For more details about Backstage permissions, see:

- [Backstage Permissions](https://backstage.io/docs/permissions/overview)
