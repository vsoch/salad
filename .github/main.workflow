workflow "Deploy ImageDefinition Schema" {
  on = "push"
  resolves = ["Extract ImageDefinition Schema"]
}

action "build" {
  uses = "actions/docker/cli@master"
  args = "build -t vanessa/salad ."
}

action "Extract ImageDefinition Schema" {
  needs = ["build"]
  uses = "docker://openschemas/extractors:ImageDefinition"
  secrets = ["GITHUB_TOKEN"]
  env = {
    IMAGE_THUMBNAIL = "https://vsoch.github.io/datasets/assets/img/avocado.png"
    IMAGE_ABOUT = "Generate ascii art for a fork or spoon, along with a pun."
    IMAGE_DESCRIPTION = "alpine base with GoLang and PUNS."
  }
  args = ["extract", 
          "--name", "vanessa/salad",
          "--contact", "@vsoch", 
          "--deploy", 
          "--filename", "/github/workspace/Dockerfile"]
}
