import unittest
import os
from pact import Verifier

os.environ['NO_PROXY'] = 'localhost,127.0.0.1,intranet,*.intranet'

class ProviderPactTest(unittest.TestCase):

    def setUp(self):
        """
        Setup the provider verification test. This includes defining
        the base URL of the provider and any headers like Authorization
        that are required.
        """
        self.provider_url = "https://awscc-eu-dev-api.barclays.intranet/bcd-dev/actions/bcd/chat-template"
        # self.auth_token = "Bearer your_dynamic_token_here"  
        self.header = ["Authorization: Bearer your_dynamic_token_here"] # Update this with the current token
    def test_verify_consumer_fullserve(self):
        """
        Verifies the contract for the 'fullserve' consumer scenario.
        """
        pact_file = "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"  # Path to the Consumer Pact file

        verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=f"{self.provider_url}?consumer=fullserve", custom_provider_headers=self.header)
        result = verifier.verify_pacts(
            pact_file
        )
        self.assertEqual(result, 0, "Provider verification failed for 'fullserve' consumer")

    def test_verify_consumer_veripark(self):
        """
        Verifies the contract for the 'veripark' consumer scenario.
        """
        pact_file = "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"  # Path to the Consumer Pact file

        verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=f"{self.provider_url}?consumer=veripark", custom_provider_headers=self.header)
        result = verifier.verify_pacts(
            pact_file
        )
        self.assertEqual(result, 0, "Provider verification failed for 'veripark' consumer")

    def test_verify_consumer_salesforce(self):
        """
        Verifies the contract for the 'salesforce' consumer scenario.
        """
        pact_file = "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"  # Path to the Consumer Pact file

        verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=f"{self.provider_url}?consumer=salesforce", custom_provider_headers=self.header)
        result = verifier.verify_pacts(
            pact_file
        )
        self.assertEqual(result, 0, "Provider verification failed for 'salesforce' consumer")

    def test_verify_consumer_invalid(self):
        """
        Verifies the contract for the 'invalid' consumer scenario.
        Expected response:
        {
            "data": {
                "type": "ChatTemplates",
                "attributes": {"templates":[]}
            }
        }
        """
        pact_file = "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"  # Path to the Consumer Pact file

        verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=f"{self.provider_url}?consumer=invalid", custom_provider_headers=self.header)
        result = verifier.verify_pacts(
            pact_file
        )
        self.assertEqual(result, 0, "Provider verification failed for 'invalid' consumer")

    def test_verify_consumer_empty(self):
        """
        Verifies the contract for the 'empty consumer' scenario.
        Expected response:
        {
            "errors":[{"title":"BadRequest", "detail": "Empty consumer id", "status": "400"}]
        }
        """
        pact_file = "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"  # Path to the Consumer Pact file

        verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=f"{self.provider_url}?consumer=", custom_provider_headers=self.header)
        result = verifier.verify_pacts(
            pact_file
        )
        self.assertEqual(result, 0, "Provider verification failed for 'empty consumer'")

    def test_verify_missing_authorization(self):
        """
        Verifies the contract when Authorization header is missing.
        Expected response:
        {
            "message": "Unauthorized"
        }
        """
        pact_file = "./pacts/toolbar-chat-api-unv-bcd-chat-pdr.json"  # Path to the Consumer Pact file

        verifier = Verifier(provider="unv-bcd-chat-pdr", provider_base_url=f"{self.provider_url}?consumer=fullserve")
        result = verifier.verify_pacts(
            pact_file
        )
        self.assertEqual(result, 0, "Provider verification failed for missing Authorization")

if __name__ == "__main__":
    unittest.main()
