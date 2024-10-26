import unittest
import os
from pact import Verifier

# Set up environment variables for the test environment
os.environ['NO_PROXY'] = 'localhost,127.0.0.1,intranet,*.intranet'

class ProviderPactTest(unittest.TestCase):

 def setUp(self):
    """
    Set up the base provider URL and configure headers required for authorization.
    """
    self.provider_base_url = "https://awscc-eu-dev-api.barclays.intranet" # Base URL only
    self.api_path = "/bcd-dev/actions/bcd/chat-template"
    self.headers = self.get_headers()

 def get_headers(self, auth_required=True):
    """Retrieve the headers with Authorization if required."""
    if auth_required:
        auth_token = os.getenv("API_AUTH_TOKEN", "Bearer your_dynamic_token_here")
        return {"Authorization": auth_token}
    return {}

 def get_pact_file_path(self):
    """Path to the Consumer Pact file."""
    return "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"

 def verify_contract(self, auth_required=True):
    """Reusable method to verify contract based on authorization need."""
    verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=self.provider_base_url + self.api_path)
    result = verifier.verify_pacts(
    self.get_pact_file_path(),
    headers=self.get_headers(auth_required),
    provider_version="1.0.0"
    )
    return result

 def test_verify_consumer_fullserve(self):
    """Verifies the contract for the 'fullserve' consumer."""
    self.assertEqual(self.verify_contract(), 0, "Provider verification failed for 'fullserve' consumer")

 def test_verify_consumer_veripark(self):
    """Verifies the contract for the 'veripark' consumer."""
    self.assertEqual(self.verify_contract(), 0, "Provider verification failed for 'veripark' consumer")

 def test_verify_consumer_salesforce(self):
    """Verifies the contract for the 'salesforce' consumer."""
    self.assertEqual(self.verify_contract(), 0, "Provider verification failed for 'salesforce' consumer")

 def test_verify_consumer_invalid(self):
    """Verifies the contract for an 'invalid' consumer."""
    self.assertEqual(self.verify_contract(), 0, "Provider verification failed for 'invalid' consumer")

 def test_verify_consumer_empty(self):
    """Verifies the contract for an empty consumer query."""
    self.assertEqual(self.verify_contract(), 0, "Provider verification failed for 'empty consumer'")

 def test_verify_missing_authorization(self):
    """Verifies the contract when Authorization header is missing."""
    self.assertEqual(self.verify_contract(auth_required=False), 0, "Provider verification failed for missing Authorization")

if __name__ == "__main__":
 unittest.main()