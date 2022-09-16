pipeline{
    agent none
    environment{
        AUTHOR = "Vincen Tampan dan Berani"
        EMAIL = "vincen@mail.com"
        WEB = "https://wwww.vincen.com"
    }

    // triggers{
    //     cron("*/5 * * * *") //every 5 minutes
    //     //pollSCM("*/5 * * * *") //every 5 minutes
    //    // upstream(upstreamProjects: 'job1,job2', threshold: hudson.model.Result.SUCCESS) //kalau job1 atau job2 SUCCESS maka pipeline ini akan jalan secara otomatis ||src: https://javadoc.jenkins-ci.org/hudson/model/Result.html
    // }

    parameters{ //name bisa dipakai jadi substitusi variable dengan ${params.NAME}
        string(name: "NAME", defaultValue: "Guest", description: "What is your name")
        text(name: "DESCRIPTION", defaultValue: "no description", description: "Tell me about yourself")
        booleanParam(name: "DEPLOY", defaultValue: false, description: "Need to Deploy")
        choice(name: "SOCIAL_MEDIA", choices: ['Instagram', 'Facebook', 'Twitter'], description: "Which Social Media")
        password(name: "SECRET", defaultValue: "", description: "Encrypt Key")
    }

    options{
        disableConcurrentBuilds() //mematikan jalan pararel job
        timeout(time: 10, unit: 'MINUTES')//totalny bukan timeout per stage
    } //BISA DILEVEL pipeline atau per Stages (ini di level pipeline)

    stages{
        //SEQUENTIAL STAGES
        stage("OS Setup"){ //START
            //matrix akan membuat stage dibuild menggunakan setiap kombinasi matrix axis valuenya (jadi bakal banyak) atau disebut Matrix Cell
            matrix{
                //matrix sifatnya seperti pararel; 1.agent ditentukan setiap stage, 2. setting failFast atau parallelAlwaysFailFast()
                axes{ //set axes nya apa sajas
                    axis{
                        name "OS" //valuenya bisa diambil dengan menggunakan name nya ${OS} (lihat pada stages.stage.steps)(VIN#2)
                        values "Windows", "Linux", "Mac"
                    }
                    axis{
                        name "ARC"
                        values "32", "64"
                    }
                }//END AXES
                //kode diatas jadinya menjadi kombinasi OS+ARC: Windows 32, Windows 64, Linux 32, Linux 64, Mac 32, Mac 64

                //INI EXCLUDE (MENGHIRAUKAN BEBERAPA MATRIX)
                excludes{
                    exclude{
                        axis{
                            name "OS"
                            values "Mac"
                        }
                        axis{
                            name "ARC"
                            values "32"
                        }
                    }
                }//END EXCLUDES 
                //berarti menghiraukan matrix yang Mac 32

                //STAGESNYA
                stages{
                    stage("OS Setup"){
                        agent{
                            node{
                                label "windows && java17"
                            }
                        }
                        steps{
                            echo("Setup ${OS} ${ARC}") //valuenya bakal diambil seperti for loop matrix(VIN#2)
                        }
                    }
                }
            }//END MATRIX
        }

        stage("Preparation"){
            //agent{node{label windows && java17}} //Kalau pakai Sequential stagesnya pakai pararel, agentnya harus diatur disetiap bracket 'pararel'
            //bagian ini harus pilih satu (biasa steps) antara :stages, pararel atau matrix
            //failFast true (VIN#1)
            parallel{ //kalau pilih parallel, semua stagenya berjalan bersamaan, formatnya seperti stages juga isinya hanya saja harus isi agent pada tiap stagenya
            //PADA PARAREL DEFAULTNYA JIKA ADA ERROR MAKA TETAP DITUNGGU (jika mau berhenti tambahkan perintah failFast true (VIN#1)) atau parallelAlwaysFailFast() di options
                stage("Prepare java"){
                    agent{
                        node{
                            label "windows && java17"
                        }
                    }
                    steps{
                        echo("Prepare Java nich boxq")
                    }
                }
                stage("Prepare Maven"){
                    agent{
                        node{
                            label "windows && java17"
                        }
                    }
                    steps{
                        echo("Prepare Maven nich boxq")
                    }
                }
            }//END
        }//END STAGE PREPARATION

        stage("Parameter"){
            agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
            steps{
                echo "Hello ${params.NAME}"
                echo "Your Description ${params.DESCRIPTION}"
                echo "Your Social media is ${params.SOCIAL_MEDIA}"
                echo "Need to deploy ${params.DEPLOY}"
                echo "Your secret is ${params.SECRET}"
            }

        }

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
            //INPUT
            input{
                message "Can we deploy?" //pesan
                ok "Yes, of course" //tulisan di tombol ok
                submitter "vincen,teman" //user yang bisa melakukan submit
                parameters{
                    choice(name: "TARGET_ENV", choices: ['DEV', 'QA', 'PROD'], description: "Which Environment?")
                } //bisa menambahkan parameter dan bisa diakses nama parameternya tapi tanpa params. karena ini di level stage 
            }
            agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
            steps{
                echo "MODE ENVIRONMENT: ${TARGET_ENV}"
                echo 'Hello Deploy 1'
                sleep(5)
                echo 'Hello Deploy 2'
                echo 'Hello Deploy 3'
            }
        }

        stage("Release"){
            when{ //stage ini jalan "jika" ||start
                expression{ //ini artinya pakai groovy syntax
                    return params.DEPLOY; //tergantung dari parameter DEPLOYnya
                }
            }//end
             agent{ //ditambah tiap stage
                node{
                    label "windows && java17"
                }
            }
            steps{
                echo "Release nih bro"
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