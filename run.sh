# Load the enviroment variables.
source .env

# Choose what private key and whether to use the mock vrf coordinator for contract deployments.
echo "PRIVATE_KEY=$PRIVATE_KEY\\nUSE_MOCK=$USE_MOCK" > packages/example-contracts/.env

# Run the prover.
if [ "$USE_MOCK" == "false" ]
then 
    pnpm concurrently -n contracts,client,prover -c cyan,magenta,blue \
        "cd packages/example-contracts && pnpm run dev" \
        "cd packages/example-client && pnpm run dev" \
        "cd packages/prover && VRF_JSON_PATH=$VRF_JSON_PATH RPC_URL=$RPC_URL PRIVATE_KEY=$PRIVATE_KEY pnpm run dev"
else
    pnpm concurrently -n contracts,client,prover -c cyan,magenta,blue \
        "cd packages/example-contracts && pnpm run dev" \
        "cd packages/example-client && pnpm run dev" \
        "cd packages/mock-prover && VRF_JSON_PATH=$VRF_JSON_PATH RPC_URL=$RPC_URL PRIVATE_KEY=$PRIVATE_KEY pnpm run dev"
fi