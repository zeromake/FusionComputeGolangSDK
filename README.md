# FusionComputeGolangSDK
Golang Client for FusionCompute 8.0

## Feature

- all func add param `ctx context.Context` http request add `SetContext`.
- use `resty SetResult` func substitute `json.Unmarshal`.
- `client.FusionComputeClient` add `SetUserType` support to account other type.
- `vm` add `ListVMVersion`, `UpdateVM` func and `ListVm` change to paging support.
