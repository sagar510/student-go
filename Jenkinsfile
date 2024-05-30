pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'develop') {
                        echo "Building develop branch"
                        // Add your build steps for develop branch here
                    } else if (env.BRANCH_NAME == 'feature') {
                        echo "Building feature branch"
                        // Add your build steps for feature branch here
                    } else {
                        echo "Building other branch"
                        // Add your build steps for other branches here
                    }
                }
            }
        }
        stage('Test') {
            steps {
                echo "Running tests for ${env.BRANCH_NAME} branch"
                // Add your test steps here
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying ${env.BRANCH_NAME} branch"
                // Add your deploy steps here
            }
        }
    }
}
