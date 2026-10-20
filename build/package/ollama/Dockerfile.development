FROM ollama/ollama:0.31.0

# setup os
RUN sed -i 's|http://archive.ubuntu.com|https://archive.ubuntu.com|g; s|http://security.ubuntu.com|https://security.ubuntu.com|g' /etc/apt/sources.list.d/ubuntu.sources \
    && apt-get update \
    && apt-get install -y --no-install-recommends curl jq \
    && rm -rf /var/lib/apt/lists/*

# copy scripts
COPY ./scripts/utilities/_set_vars.sh /_set_vars.sh
COPY ./build/package/ollama/_cleanup.bash /_cleanup.bash
COPY ./build/package/ollama/entrypoint.bash /entrypoint.bash

RUN chmod +x /_set_vars.sh /_cleanup.bash /entrypoint.bash

ENTRYPOINT ["/entrypoint.bash"]
