ARG BUILDER_CONTAINER_IMAGE

FROM ${BUILDER_CONTAINER_IMAGE} AS build

FROM scratch

COPY --from=build /var/workspace/grm.exe .

ENTRYPOINT [ "/grm.exe" ]

VOLUME [ "/var/workspace" ]

WORKDIR /var/workspace
