pipeline {
    agent any
    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows'], description: 'Pick OS')
        choice(name: 'TARGETARCH', choices: ['amd64', 'arm64'], description: 'Pick architecture')
    }
    environment{
        REPO = 'https://github.com/bartaadalbert/kbot'
        BRANCH = 'main'
    }
    stages {
        stage('clone') {
            steps {
                echo 'Clone repo'
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }
        stage('test') {
            steps {
                echo 'run test'
                sh 'make test'
            }
        }
        stage('build') {
            steps {
                echo 'Build'
                sh 'make build'
            }
        }
        stage('image') {
            steps {
                echo 'Create docker image'
                sh 'make image'
            }
        }
        stage('push') {
            steps {
                script{
                    docker.withRegistry('','dockerhub'){
                        sh 'make push'
                    }
                }
            }
        }
    }
}
