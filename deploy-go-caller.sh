cd go/project_caller

gcloud builds submit --tag=eu.gcr.io/project-booster-perf-disp/projectcaller:latest .

gcloud beta run deploy --image eu.gcr.io/project-booster-perf-disp/projectcaller:latest --platform managed --region europe-west1 --allow-unauthenticated projectcaller --set-env-vars PATH_HELLO_WORLD=https://projectemitter-zlg45lmdqa-ew.a.run.app/