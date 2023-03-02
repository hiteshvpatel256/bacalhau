# coding: utf-8

"""
    Bacalhau API

    This page is the reference of the Bacalhau REST API. Project docs are available at https://docs.bacalhau.org/. Find more information about Bacalhau at https://github.com/bacalhau-project/bacalhau.  # noqa: E501

    OpenAPI spec version: 0.3.23.post7
    Contact: team@bacalhau.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""

import pprint
import re  # noqa: F401

import six


class SubmitRequest(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        "client_public_key": "str",
        "job_create_payload": "list[int]",
        "signature": "str",
    }

    attribute_map = {
        "client_public_key": "client_public_key",
        "job_create_payload": "job_create_payload",
        "signature": "signature",
    }

    def __init__(
        self, client_public_key=None, job_create_payload=None, signature=None
    ):  # noqa: E501
        """SubmitRequest - a model defined in Swagger"""  # noqa: E501
        self._client_public_key = None
        self._job_create_payload = None
        self._signature = None
        self.discriminator = None
        self.client_public_key = client_public_key
        self.job_create_payload = job_create_payload
        self.signature = signature

    @property
    def client_public_key(self):
        """Gets the client_public_key of this SubmitRequest.  # noqa: E501

        The base64-encoded public key of the client:  # noqa: E501

        :return: The client_public_key of this SubmitRequest.  # noqa: E501
        :rtype: str
        """
        return self._client_public_key

    @client_public_key.setter
    def client_public_key(self, client_public_key):
        """Sets the client_public_key of this SubmitRequest.

        The base64-encoded public key of the client:  # noqa: E501

        :param client_public_key: The client_public_key of this SubmitRequest.  # noqa: E501
        :type: str
        """
        if client_public_key is None:
            raise ValueError(
                "Invalid value for `client_public_key`, must not be `None`"
            )  # noqa: E501

        self._client_public_key = client_public_key

    @property
    def job_create_payload(self):
        """Gets the job_create_payload of this SubmitRequest.  # noqa: E501

        The data needed to submit and run a job on the network:  # noqa: E501

        :return: The job_create_payload of this SubmitRequest.  # noqa: E501
        :rtype: list[int]
        """
        return self._job_create_payload

    @job_create_payload.setter
    def job_create_payload(self, job_create_payload):
        """Sets the job_create_payload of this SubmitRequest.

        The data needed to submit and run a job on the network:  # noqa: E501

        :param job_create_payload: The job_create_payload of this SubmitRequest.  # noqa: E501
        :type: list[int]
        """
        if job_create_payload is None:
            raise ValueError(
                "Invalid value for `job_create_payload`, must not be `None`"
            )  # noqa: E501

        self._job_create_payload = job_create_payload

    @property
    def signature(self):
        """Gets the signature of this SubmitRequest.  # noqa: E501

        A base64-encoded signature of the data, signed by the client:  # noqa: E501

        :return: The signature of this SubmitRequest.  # noqa: E501
        :rtype: str
        """
        return self._signature

    @signature.setter
    def signature(self, signature):
        """Sets the signature of this SubmitRequest.

        A base64-encoded signature of the data, signed by the client:  # noqa: E501

        :param signature: The signature of this SubmitRequest.  # noqa: E501
        :type: str
        """
        if signature is None:
            raise ValueError(
                "Invalid value for `signature`, must not be `None`"
            )  # noqa: E501

        self._signature = signature

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(
                    map(lambda x: x.to_dict() if hasattr(x, "to_dict") else x, value)
                )
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(
                    map(
                        lambda item: (item[0], item[1].to_dict())
                        if hasattr(item[1], "to_dict")
                        else item,
                        value.items(),
                    )
                )
            else:
                result[attr] = value
        if issubclass(SubmitRequest, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, SubmitRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
