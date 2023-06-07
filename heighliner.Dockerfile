ARG BASE_IMAGE_TAG=latest

FROM ghcr.io/strangelove-ventures/infra-toolkit:v0.0.7 AS infra-toolkit

FROM ghcr.io/duality-labs/duality:${BASE_IMAGE_TAG}

# NOTE: Make heighliner owner of all copied files so
# they can be deleted after startup if desired

# Copy all binaries to /usr/bin
COPY --from=infra-toolkit --chown=heighliner /usr/bin /usr/bin
COPY --from=infra-toolkit --chown=heighliner /usr/local/bin /usr/bin
COPY --from=infra-toolkit --chown=heighliner /bin/ /usr/bin/
COPY --from=infra-toolkit --chown=heighliner /sbin/ /usr/bin/

# Copy busybox to /bin since most binaries are symlinked to it
COPY --from=infra-toolkit --chown=heighliner /bin/busybox /bin/

# Copy in required libraries
COPY --from=infra-toolkit --chown=heighliner /lib /lib/
COPY --from=infra-toolkit --chown=heighliner /usr/lib /usr/lib/

COPY scripts scripts
COPY networks networks

CMD ["sh", "./scripts/startup.sh"]