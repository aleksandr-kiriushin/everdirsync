pipeline {
    agent {
        docker {
            image 'golang:1.16-alpine'
            args '-e XDG_CACHE_HOME="/tmp/.cache"'
        }
    }
    stages {
        stage('Build') {
            steps {
                sh 'go build'
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
                GITHUB_TOKEN = credentials('github-jenkins-releases')
            }
            steps {
                sh 'curl -sL https://git.io/goreleaser | bash'
            }
        }
    }
}
