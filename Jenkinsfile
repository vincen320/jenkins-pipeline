pipeline{
    agent{
        node{
            label "windows && java17"
        }
    }

    stages{
        stage('Build'){
            steps{
                echo 'Hello Build'
            }
        }

        stage('Test'){
            steps{
                echo 'Hello Test'
            }
        }

        stage('Deploy'){
            steps{
                echo 'Hello Deploy'
            }
        }
    }
    post{
        always{
            echo 'I will always say Hello again!'
        }
        success{
            echo 'Yay, Success'
        }
        failure{
            echo 'Oh no, failure'
        }
        cleanup{
            echo "Dont't care success or not"
        }
    }
}