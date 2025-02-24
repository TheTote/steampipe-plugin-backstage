FROM ghcr.io/turbot/steampipe

# Setup prerequisites (as root)
USER root:0
RUN apt-get update -y \
    && apt-get install -y git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create plugin directory and set ownership
RUN mkdir -p /home/steampipe/.steampipe/plugins/local/backstage

# Copy local config and plugin
COPY .steampipe/config/backstage.spc /home/steampipe/.steampipe/config/backstage.spc
COPY build/steampipe-plugin-backstage.plugin /home/steampipe/.steampipe/plugins/local/backstage/backstage.plugin

# Set permissions while still root
RUN chmod 755 /home/steampipe/.steampipe/plugins/local/backstage/backstage.plugin \
    && chown -R steampipe:0 /home/steampipe/.steampipe

# Switch to steampipe user after setting up permissions
USER steampipe:0 