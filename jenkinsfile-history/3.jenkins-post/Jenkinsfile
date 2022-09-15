pipeline{
    agent{
        node{
            label "windows && java17"
        }
    }
    stages{
        stage('Hello'){
            steps{
                echo 'Hello World'
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