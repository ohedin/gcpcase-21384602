cd js/api

gcloud builds submit --tag gcr.io/$1/jsapi

gcloud run deploy --image gcr.io/$1/jsapi --platform managed --region europe-west1 --allow-unauthenticated jsapi