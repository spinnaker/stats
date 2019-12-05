# Spinnaker Stats Service

This is the service that collects telemetry data from customer Spinnaker instances.

This data helps the project maintainers and community decide where to focus our efforts and gauge how the product is being used in the wild.

You can find aggregations and analysis of the collected data [here](TODO). For access to the raw data, you should be a contributing member of the community, and can request access by ... TODO.



## Development

All commits to the `master` branch are deployed to a staging environment hosted at https://stats-staging.spinnaker.io. See [cloudbuild-deploy-staging.yaml](./cloudbuild-deploy-staging.yaml).

Deploying to production simply promotes the staging container to the production environment. See [cloudbuild-promote-staging-to-prod.yaml](cloudbuild-promote-staging-to-prod.yaml).
