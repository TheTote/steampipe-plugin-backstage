# Backstage plugin for Steampipe

Use SQL to query namespaces, rules and more from [Backstage](https://backstage.io/).

```sql
select
  name,
  kind,
  namespace,
  description
from
  backstage_catalog_entity
where 
  kind = 'Component';
```

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables.md)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```bash
steampipe plugin install chussenot/backstage
```

## Development

To build the plugin and install it in your `.steampipe` directory

```bash
make
```

Copy the default config file:

```bash
make config
```

## License

Apache 2

## Resources

- [steampipe](https://steampipe.io)
- [backstage](https://backstage.io/)
- [plugin release checklist](https://steampipe.io/docs/develop/plugin-release-checklist)
- [go-backstage](https://github.com/datolabs-io/go-backstage)
- [backstage-terraform](https://registry.terraform.io/providers/datolabs-io/backstage/latest/docs)
- [steampipe plugin standards](https://steampipe.io/docs/develop/standards#naming)