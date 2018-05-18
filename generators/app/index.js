'use strict';
const Generator = require('yeoman-generator');
const chalk = require('chalk');
const yosay = require('yosay');
const path = require('path');
const mkdir = require('mkdirp');

module.exports = class extends Generator {
  paths() {
    this.destinationRoot(process.env.GOPATH || './');
  }

  prompting() {
    this.log(
      yosay('Welcome to the sensational ' + chalk.red('generator-go') + ' generator!')
    );

    const prompts = [
      {
        type: 'input',
        name: 'projectOwner',
        message: 'What is the project owner?',
        required: true
      },
      {
        type: 'input',
        name: 'projectName',
        message: 'What is the project name?',
        required: true
      },
      {
        type: 'input',
        name: 'port',
        message: 'What port will your app run?',
        default: 8080,
        required: true
      }
    ];

    return this.prompt(prompts).then(props => {
      // To access props later use this.props.someAnswer;
      this.props = props;
    });
  }

  default() {
    this.props.projectRoot = path.join(
      'github.com/',
      this.props.projectOwner,
      this.props.projectName
    );

    console.log('Generating tree folders');
    this.pkgDir = this.destinationPath('pkg');
    this.srcDir = this.destinationPath(path.join('src/', this.props.projectRoot));
    this.binDir = this.destinationPath('bin');

    mkdir.sync(this.pkgDir);
    mkdir.sync(this.srcDir);
    mkdir.sync(this.binDir);
  }

  writing() {
    let ctx = {
      projectName: this.props.projectName,
      projectRoot: this.props.projectRoot,
      port: this.props.port
    };

    this.fs.copyTpl(this.templatePath('**/*'), this.srcDir, ctx);
    this.fs.copy(this.templatePath('**/.*'), this.srcDir);
  }

  end() {
    this.log('Thanks for using our Go Generator.');
  }
};
