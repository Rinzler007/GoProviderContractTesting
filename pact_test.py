import os
import unittest
from pact import Verifier

class ProviderPactTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        """Initialize paths and configurations."""
        # Base URL where the provider is running (without endpoint and query parameters)
        cls.provider_base_url = 'http://localhost:8080'

        # Provider and Consumer names
        cls.provider_name = 'unv-bcd-chat-pdr'
        cls.consumer_name = 'toolbar-chat-api'

        # Path to the pact file directory
        cls.pact_dir = os.path.abspath(os.path.join(os.path.dirname(__file__), 'pacts'))
        
        # Ensure the pact directory exists
        if not os.path.exists(cls.pact_dir):
            raise FileNotFoundError(f"Pact directory not found at {cls.pact_dir}")

    def test_provider_pacts(self):
        """Verify provider against the pact file."""
        verifier = Verifier(
            provider=self.provider_name,
            provider_base_url=self.provider_base_url,
            state_change_url='http://localhost:8080/setup'  # Endpoint to set provider states
            # Uncomment and set if using a Pact Broker
            # pact_broker_url=os.getenv('PACT_BROKER_URL'),
            # broker_username=os.getenv('PACT_BROKER_USERNAME'),
            # broker_password=os.getenv('PACT_BROKER_PASSWORD'),
        )

        # Path to the specific pact file
        pact_file_relative = 'toolbar-chat-api-unv-bcd-chat-pdr.json'
        pact_file_absolute = os.path.join(self.pact_dir, pact_file_relative)

        # Check if the pact file exists
        if not os.path.isfile(pact_file_absolute):
            raise FileNotFoundError(f"Pact file not found at {pact_file_absolute}")

        # Create the file URL
        pact_url = f'file://{pact_file_absolute}'

        # Verify the pact
        verification_result = verifier.verify_pacts(
            provider=self.provider_name,
            pact_urls=[pact_url],
            publish_verification_results=False  # Set to True if publishing to Pact Broker
        )

        # Assert verification success
        self.assertTrue(verification_result, "Pact verification failed.")

if __name__ == '__main__':
    unittest.main()
