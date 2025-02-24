# Table: backstage_catalog_template

A Template in Backstage is a blueprint that can be used to create new software components, documentation sites, and other resources. Templates help standardize creation of new entities in your organization.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the template. |
| namespace | string | The namespace the template belongs to. |
| kind | string | The kind of the entity (always "Template"). |
| metadata | json | The full metadata of the template. |
| spec | json | The specification data of the template. |
| relations | json | The relations of the template to other entities. |
| title | string | A display name of the template. |
| description | string | A description of the template. |
| labels | json | Labels attached to the template. |
| annotations | json | Annotations attached to the template. |
| tags | json | A list of tags attached to the template. |
| links | json | A list of external hyperlinks related to the template. |
| type | string | Type of the template. |
| parameters | json | Parameters defined in the template. |
| steps | json | Steps defined in the template. |

## Examples

### Basic info

```sql
select
  name,
  type,
  description
from
  backstage_catalog_template;
```

### List templates with their parameters

```sql
select
  name,
  parameters
from
  backstage_catalog_template
where
  parameters is not null;
```

### Find templates by type

```sql
select
  type,
  count(*) as template_count
from
  backstage_catalog_template
group by
  type
order by
  template_count desc;
``` 
