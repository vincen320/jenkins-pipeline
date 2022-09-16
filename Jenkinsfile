pipeline{
    agent none
    environment{
        AUTHOR = "Vincen Tampan dan Berani"
        EMAIL = "vincen@mail.com"
        WEB = "https://wwww.vincen.com"
    }

    options{
        disableConcurrentBuild() //mematikan jalan pararel job
        timeout(time: 10, unit: 'MINUTES')
    } //BISA DILEVEL pipeline atau per Stages (ini di level pipeline)

    stages{
        stage('Prepare'){
            environment{ //bebas environmentnya mau disini atau global, kalau dinsi berarti cuma bisa dipakai di bracket ini aja (stage Prepare)
                NAMABEBAS = credentials("vincen_rahasia") //pakai id credentials || terbentuk 2 varible yaitu NAMABEBAS_USR & NAMABEBAS_PSW
            }
            agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
            steps{
                echo("Author ${AUTHOR}")
                echo("Email ${EMAIL}")
                echo("Web ${WEB}")
                echo("Start Job: ${env.JOB_NAME}")
                echo("Start Build: ${env.BUILD_NUMBER}")
                echo("Branch Name: ${env.BRANCH_NAME}")
                echo("App User: ${NAMABEBAS_USR}")
                echo("App Password: ${NAMABEBAS_PSW}")
                bat("echo 'App Password: ${NAMABEBAS_PSW}' > rahasia.txt") //ini tidak aman
                //Cara aman untuk data sensitive spt password(?)
                bat('echo "App Password: $NAMABEBAS_PSW" > "rahasia2.txt"') // ini aman (pakai tanda petik satu)
            }
        }

        stage('Build'){
            agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
            steps{
                script{ //start script || script Groovy || harus dibuat dalam 'steps'
                    for (i = 0; i < 10 ; i++){
                        echo("Script ${i}")
                    }
                }//end script
                echo 'Start Build'
                //bat('go build -o user-service main.go') //mahal resource ssd
                echo 'Finish Build'
            }
        }

        stage('Test'){
            agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
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
            agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
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

//Options
//https://www.jenkins.io/doc/book/pipeline/syntax/#options