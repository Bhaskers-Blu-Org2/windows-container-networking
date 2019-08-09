pipeline {
    agent any
    options {
        skipDefaultCheckout true
    }
    stages {
        stage('Checkout') {
            steps {
                dir('src/github.com/Microsoft/windows-container-networking') {
                    checkout scm
                }
            }
        }
        stage('Build') {
            steps {
               bat '''
               set GO111MODULE=on
               set GOPATH=%cd%;C:\\users\\nathan\\go
               cd src/github.com/Microsoft/windows-container-networking
               make vendor
               make all
               '''
           }
        }
        stage('Test') {
             steps {
               bat '''
               set GOPATH=%cd%
               go env
               cd src/github.com/Microsoft/windows-container-networking
               make test
               '''
            }
        }
    }
}