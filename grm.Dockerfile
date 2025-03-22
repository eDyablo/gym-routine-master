ARG BUILDER_CONTAINER_IMAGE
ARG SOURCE_DIR=grm
ARG BINARY_OUTPUT_PATH=/var/workspace/grm.exe
FROM ${BUILDER_CONTAINER_IMAGE} AS build

FROM scratch

COPY --from=build /var/workspace/grm.exe .

ENTRYPOINT [ "/grm.exe" ]

VOLUME [ "/var/workspace" ]

WORKDIR /var/workspace
