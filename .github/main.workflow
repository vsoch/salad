workflow "Deploy ImageDefinition Schema" {
  on = "push"
  resolves = ["Extract ImageDefinition Schema"]
}

action "Extract ImageDefinition Schema" {
  uses = "openschemas/extractors/ImageDefinition@master"
  secrets = ["GITHUB_TOKEN"]
  env = {
    IMAGE_THUMBNAIL = "https://vsoch.github.io/datasets/assets/img/avocado.png"
    IMAGE_ABOUT = "Generate ascii art for a fork or spoon, along with a pun."
    IMAGE_DESCRIPTION = "alpine base with GoLang and PUNS."
  }
  args = ["extract", "--contact", "@vsoch", '--deploy']
}

