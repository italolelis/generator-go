FROM golang:1-alpine AS builder

WORKDIR /go/src/<%=projectRoot%>

COPY . ./

RUN apk add --update bash make git
RUN make

# ---

FROM mattjtodd/raml2html AS ramlbuilder

COPY . ./

RUN raml2html -i docs/specification/api.raml -o "/api.html"

# ---

FROM alpine

COPY --from=ramlbuilder /api.html /docs/specification/api.html
COPY --from=builder /go/src/<%=projectRoot%>/dist/<%=projectName%> /

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

HEALTHCHECK --interval=5s --timeout=5s --retries=3 CMD curl -f http://localhost:8081/status || exit 1

EXPOSE <%=port%>
ENTRYPOINT ["/<%=projectName%>"]
