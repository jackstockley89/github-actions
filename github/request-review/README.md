## Example
---
```
name: Assign Collaborator to Pull Request

on:
  pull_request:
  workflow_dispatch:

jobs:
  assignee:
    runs-on: ${{ matrix.os }}

    strategy:
      matrix:
        os: [ubuntu-latest]

    steps:
      - name: Checkout
        uses: actions/checkout@v3
      
      - name: adding collaborator
        uses: docker://jackstock8904/assignee:latest
        if: steps.approve_colab_user.outputs.approve_colab_user == 'true'
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```