## Example
---
```
name: Pull Request Authentication Check

on:
  pull_request:
  workflow_dispatch:

jobs:
  check-pr-user:
    runs-on: ${{ matrix.os }}

    strategy:
      matrix:
        os: [ubuntu-latest]

    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - id: approve_colab_user
        name: Check PR user 
        uses: docker://jackstock8904/auth-check:latest
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

      - name: Output
        run: echo ${{ steps.approve_colab_user.outputs.approve_colab_user }}

      - name: Comment on PR if true
        uses: actions/github-script@v6
        if: steps.approve_colab_user.outputs.approve_colab_user == 'true'
        with:
          github-token: "${{ secrets.GITHUB_TOKEN }}"
          script: |
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: 'This is a known Colaborator. Please Review this Pull Request.'
            });
      
      - name: Comment on PR if false
        uses: actions/github-script@v6
        if: steps.approve_colab_user.outputs.approve_colab_user == 'false'
        with:
          github-token: "${{ secrets.GITHUB_TOKEN }}"
          script: |
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: 'This isn't a known Colaborator. This Pull Request will now be closed.'
            });
```