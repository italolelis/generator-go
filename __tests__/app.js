'use strict';
const path = require('path');
const assert = require('yeoman-assert');
const helpers = require('yeoman-test');

describe('generator-go:app', () => {
  let srcDir = path.join(process.env.GOPATH, 'src/github.com/hellofresh/my-test');

  beforeAll(() => {
    return helpers
      .run(path.join(__dirname, '../generators/app'))
      .withPrompts({ projectName: 'my-test', port: 8080 });
  });

  it('creates files', () => {
    assert.file([
      path.join(srcDir, 'cmd/root.go'),
      path.join(srcDir, 'cmd/server_start.go'),
      path.join(srcDir, 'cmd/version.go'),
      path.join(srcDir, 'docs/specification/types/Status.raml'),
      path.join(srcDir, 'docs/specification/api.raml'),
      path.join(srcDir, 'pkg/config/config.go'),
      path.join(srcDir, 'pkg/handler/docs.go'),
      path.join(srcDir, 'pkg/handler/github.go'),
      path.join(srcDir, 'pkg/handler/hello_world.go'),
      path.join(srcDir, 'pkg/metrics/metrics.go'),
      path.join(srcDir, 'pkg/tracer/jaeger_log.go'),
      path.join(srcDir, 'pkg/tracer/middleware.go'),
      path.join(srcDir, 'pkg/tracer/tracer.go'),
      path.join(srcDir, '.editorconfig'),
      path.join(srcDir, '.gitignore'),
      path.join(srcDir, 'Dockerfile'),
      path.join(srcDir, 'Gopkg.toml'),
      path.join(srcDir, 'main.go'),
      path.join(srcDir, 'Makefile')
    ]);
  });
});
