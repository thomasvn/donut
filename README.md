# 🍩 Donut

A CLI tool modeled after <https://www.donut.com>

- track friends you'd like to keep in touch with
- get randomly paired with a friend to reach out to

## Install

```bash
$ git clone git@github.com:thomasvn/donut.git
$ cd donut
$ go build -o donut
```

## Example

```bash
$ ./donut list
Friends list:
- Alice (met on [2023-01-01 2023-02-15])
- Bob (met on [2023-01-10])
- Charlie (met on [])

$ ./donut pair
🍩 Reach out to: Alice (last met: 2023-02-15)? Accept (y), Decline (n): n
🍩 Reach out to: Bob (last met: 2023-01-10)? Accept (y), Decline (n): y
You have chosen to reach out to: Bob

# Using a custom filename
$ ./donut family.yaml list
Friends list:
- Fluffy (met on [2023-03-12 2023-04-30])
- Whiskers (met on [2023-05-01])
- Paws (met on [])
- Bubbles (met on [2023-06-20 2023-07-15 2023-08-10])
- Sparky (met on [2023-09-01])
- Luna (met on [])
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
