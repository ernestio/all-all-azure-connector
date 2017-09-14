# ALL-ALL-AZURE-CONNECTOR
Master - [![CircleCI](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/master.svg?style=shield)](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/master)    
Develop - [![CircleCI](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/develop.svg?style=shield)](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/develop)


master - [![CircleCI](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/master.svg?style=shield)](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/master)

develop - [![CircleCI](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/develop.svg?style=shield)](https://circleci.com/gh/ernestio/all-all-azure-connector/tree/develop)


Service to manage azure component actions, you can configure it to respond to different component actions setting COMPONENTS environment variable
```
$ COMPONENTS=nat.create.azure,nat.update.azure,nat.delete.azure all-all-azure-connector
```

And responds respectively with original_subject.error or original_subjet.done respectively

## Installation

```
make deps
make install
```

## Running Tests

```
make test
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).

