pipeline {
    agent {
        docker {
            image 'golang:1.16'
            args '-e XDG_CACHE_HOME="/tmp/.cache"'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build'
                sh 'rm ./everdirsync'
            }
        }
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
        stage ('Release') {
            when {
                buildingTag()
            }
            environment {
                GITHUB_TOKEN = credentials('everdirsync')
            }
            steps {
                sh 'apt update'
                sh 'apt install -y git'
                sh 'curl -sL https://git.io/goreleaser | bash'
            }
        }
    }
}
