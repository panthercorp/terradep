package statics

const TerradepConfig = `repositories:
  git:
  - url: https://github.com
    username: username
    password: yourPassword
    sshPrivateKeysPath: ~/.ssh/id_rsa
    token: abcdef1234567890
    alias: git-github
  - url: https://gitlab.com
    username: username
    password: yourPassword
    token: abcdef1234567890
    sshPrivateKeysPath: ~/.ssh/id_rsa
    alias: git-gitlab
  http:
  - url: https://yourrepository.com
    method: GET
    username: username
    password: yourPassword
    headers:
    - key: User-Agent
      value : Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36
    - key: Content-Type
      value: application/json
    alias: http-yourrepository
`

const DcfConfig = `metadata:
  name: moduleX
  version: 1.0
dependencies:
- name: moduleY
  version: 1.0
  use: git-github
  branch: master
  repo: username/repository
  subdir: subdirectory
  type: public
- name: moduleZ
  version: 1.0
  use: git-gitlab
  branch: master
  repo: username/repository
  subdir: subdirectory
  type: private
`
