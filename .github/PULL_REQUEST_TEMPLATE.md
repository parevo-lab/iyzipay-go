# Pull Request

## Description
Please include a summary of the changes and the related issue. Please also include relevant motivation and context.

Fixes # (issue)

## Type of change
Please delete options that are not relevant.

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Performance improvement
- [ ] Refactoring (no functional changes)
- [ ] Test coverage improvement

## How Has This Been Tested?
Please describe the tests that you ran to verify your changes.

- [ ] Unit tests
- [ ] Integration tests
- [ ] Manual testing with sandbox environment
- [ ] Manual testing with production environment (if applicable)

## Test Configuration:
* Go version:
* OS:
* Test environment (sandbox/production):

## Checklist:
- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published

## Breaking Changes
If this is a breaking change, please describe what changes users need to make:

```go
// Old API
client.OldMethod(param1, param2)

// New API  
client.NewMethod(&Request{
    Param1: param1,
    Param2: param2,
})
```

## Additional Notes
Add any other notes about the pull request here.