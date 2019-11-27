cd go/project_emitter

gcloud builds submit --tag=eu.gcr.io/project-booster-perf-disp/projectemitter:latest .

gcloud beta run deploy --image eu.gcr.io/project-booster-perf-disp/projectemitter:latest --platform managed --region europe-west1 --allow-unauthenticated projectemitter