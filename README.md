# 🍩 Donut

A CLI tool modeled after <https://www.donut.com>

- track friends you'd like to keep in touch with
- get randomly paired with a friend to reach out to

## Install

```bash
$ git clone
$ go build -o donut
```

## Example

Using the default `friends.yaml` provided in the repo:

```bash
./donut list
./donut pair 
```

Using your own friend files:

```bash
./donut -f "family.yaml" list
./donut -f "family.yaml" pair
```

<!-- TODO:
- Add in a "-h" option to the CLI
- The pair function should place greater weight for those who you haven't met with as much?
- BubbleTea UI
- Repo: Dependabot
-->

<!-- DONE:
- Don't add friends through the CLI, add it through a YAML file instead
- Maintain a list of people you've paired with.
- Option to accept/decline the person you were paired with. And show when you last met with the person.
- Option to pass in a custom "friends.yaml" filename. You can have one YAML per friend group.
- Add in a "TMP" date once you've accepted a donut
-->