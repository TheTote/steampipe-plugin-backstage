# Table Reference

This document provides details on the tables available in the Backstage plugin for Steampipe.

## Available Tables

| Table Name | Description |
| --- | --- |
| backstage_catalog_entity | Generic Backstage entity table |
| backstage_catalog_component | Backstage component table |
| backstage_catalog_template | Backstage template table |
| backstage_catalog_api | Backstage API table |
| backstage_catalog_group | Backstage group table |
| backstage_catalog_user | Backstage user table |
| backstage_catalog_resource | Backstage resource table |
| backstage_catalog_system | Backstage system table |
| backstage_catalog_domain | Backstage domain table |
| backstage_catalog_location | Backstage location table |

### backstage_catalog_component

Query software components in your Backstage catalog.

#### Examples

List all components with their APIs:

```sql
select
  c.name as component_name,
  c.description,
  a.name as api_name,
  a.spec ->> 'type' as api_type
from
  backstage_catalog_component as c
  left join backstage_catalog_api as a on a.spec ->> 'owner' = c.name;
```
