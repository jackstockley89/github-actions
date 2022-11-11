## Example
---
```
name: Close Pull Request

on:
  pull_request:
  workflow_dispatch:

jobs:
  close-pr:
    runs-on: ${{ matrix.os }}

    strategy:
      matrix:
        os: [ubuntu-latest]

    steps:
      - name: Checkout
        uses: actions/checkout@v3
      
      - name: close pull request
        uses: docker://jackstock8904/close-pr:latest
        if: steps.approve_colab_user.outputs.approve_colab_user == 'false'
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```