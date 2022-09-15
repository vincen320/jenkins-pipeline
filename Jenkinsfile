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

        stage('Simuasi Error'){
            steps{
                sh("stage ini error, jadi stage dibawah juga error (tidak dijalankan) (bisa lihat di pipeline stage view")
                sh("kalau mau berhasil coba komen untuk stage ini")
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