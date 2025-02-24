FROM ghcr.io/turbot/steampipe

# Setup prerequisites (as root)
USER root:0
RUN apt-get update -y \
    && apt-get install -y git \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Create all necessary directories with proper ownership
RUN mkdir -p /home/steampipe/.steampipe/plugins/local/backstage \
    && mkdir -p /home/steampipe/.steampipe/config \
    && mkdir -p /home/steampipe/.steampipe/db \
    && mkdir -p /home/steampipe/.steampipe/logs

# Copy local config and plugin
COPY .steampipe/config/backstage.spc /home/steampipe/.steampipe/config/backstage.spc
COPY build/steampipe-plugin-backstage.plugin /home/steampipe/.steampipe/plugins/local/backstage/backstage.plugin

# Set permissions and ownership for the entire .steampipe directory
RUN chmod 755 /home/steampipe/.steampipe/plugins/local/backstage/backstage.plugin \
    && chown -R steampipe:0 /home/steampipe/.steampipe \
    && chmod -R 755 /home/steampipe/.steampipe

# Switch to steampipe user after setting up permissions
USER steampipe:0 