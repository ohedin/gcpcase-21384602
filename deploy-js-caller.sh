cd js/caller

gcloud builds submit --tag gcr.io/$1/jscaller

gcloud run deploy --image gcr.io/$1/jscaller --platform managed --region europe-west1 --allow-unauthenticated jscaller