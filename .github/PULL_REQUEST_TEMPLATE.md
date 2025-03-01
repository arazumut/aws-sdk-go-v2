# **PLEASE READ BEFORE CONTINUING**

Do not submit pull requests that directly modify generated source files, e.g. `/service/s3/api_client.go`. Generated source files will always include an identifying header:

```
// Code generated by smithy-go-codegen DO NOT EDIT.
```

Manual changes to these files will be overwritten by code generation that occurs as part of the daily SDK release process.

Do not submit pull requests that directly modify files in the `/codegen/sdk-codegen/aws-models` folder. These are API model files, owned by each AWS service team, that are updated automatically as part of the daily SDK release process. Local changes to these files will not persist.

If you believe the contents of any of these files need to be changed, please [open an issue](https://github.com/aws/aws-sdk-go-v2/issues/new/choose).

#

If the PR addresses an existing bug or feature, please reference it here.

To help speed up the process and reduce the time to merge please ensure that `Allow edits by maintainers` is checked before submitting your PR. This will allow the project maintainers to make minor adjustments or improvements to the submitted PR, allow us to reduce the roundtrip time for merging your request.

# Changelog

Make sure that your pull request contains a changelog entry. You can run `make external-changelog` to create a placeholder entry. See [Changelog documents in CONTRIBUTING.md](aws/aws-sdk-go-v2/blob/main/CONTRIBUTING.md#changelog-documents) for more details.

You can generate a UUID either by running `uuidgen` or by visiting a site like [uuidtools](https://www.uuidtools.com/v4)