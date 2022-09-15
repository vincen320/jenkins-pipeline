pipeline{
    agent{
        node{
            label "windows && java17"
        }
    }

    stages{
        stage('Build'){
            steps{
                script{ //start script || script Groovy || harus dibuat dalam 'steps'
                    for (i = 0; i < 10 ; i++){
                        echo("Script ${i}")
                    }
                }//end script
                echo 'Start Build'
                bat('go build -o user-service main.go')
                echo 'Finish Build'
            }
        }

        stage('Test'){
            
            steps{
                script{ //start script || script Groovy || harus dibuat dalam 'steps'
                def data = [
                    "firstName": "Vincen",
                    "lastName": "Tampan"
                ]
                writeJSON(file: "data.json", json: data) //Ini Plugin Utility Stepsnya
                }//end script

                echo 'Start Test'
                bat('go test github.com/vincen320/user-service/service -cover')
                echo 'End Test'
            }
        }

        stage('Deploy'){
            steps{
                echo 'Hello Deploy 1'
                sleep(5)
                echo 'Hello Deploy 2'
                echo 'Hello Deploy 3'
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

//Pieline Utility Steps
//https://www.jenkins.io/doc/pipeline/steps/pipeline-utility-steps/

//Install plugin ini dulu
//https://plugins.jenkins.io/pipeline-utility-steps/