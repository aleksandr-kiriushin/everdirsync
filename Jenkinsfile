pipeline {
    agent {
        docker {
            image 'golang:1.16-alpine'
            args '-e XDG_CACHE_HOME="/tmp/.cache"'
        }
    }
    stages {
        stage('Test') {
            steps {
                sh 'go version'
                sh 'go test'
            }
        }
    }
}
