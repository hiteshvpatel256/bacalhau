# coding: utf-8

"""
    Bacalhau API

    This page is the reference of the Bacalhau REST API. Project docs are available at https://docs.bacalhau.org/. Find more information about Bacalhau at https://github.com/bacalhau-project/bacalhau.  # noqa: E501

    OpenAPI spec version: ${VERSION}
    Contact: team@bacalhau.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""

import pprint
import re  # noqa: F401

import six

class BuildVersionInfo(object):
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
        'build_date': 'str',
        'goarch': 'str',
        'goos': 'str',
        'git_commit': 'str',
        'git_version': 'str',
        'major': 'str',
        'minor': 'str'
    }

    attribute_map = {
        'build_date': 'BuildDate',
        'goarch': 'GOARCH',
        'goos': 'GOOS',
        'git_commit': 'GitCommit',
        'git_version': 'GitVersion',
        'major': 'Major',
        'minor': 'Minor'
    }

    def __init__(self, build_date=None, goarch=None, goos=None, git_commit=None, git_version=None, major=None, minor=None):  # noqa: E501
        """BuildVersionInfo - a model defined in Swagger"""  # noqa: E501
        self._build_date = None
        self._goarch = None
        self._goos = None
        self._git_commit = None
        self._git_version = None
        self._major = None
        self._minor = None
        self.discriminator = None
        if build_date is not None:
            self.build_date = build_date
        if goarch is not None:
            self.goarch = goarch
        if goos is not None:
            self.goos = goos
        if git_commit is not None:
            self.git_commit = git_commit
        if git_version is not None:
            self.git_version = git_version
        if major is not None:
            self.major = major
        if minor is not None:
            self.minor = minor

    @property
    def build_date(self):
        """Gets the build_date of this BuildVersionInfo.  # noqa: E501


        :return: The build_date of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._build_date

    @build_date.setter
    def build_date(self, build_date):
        """Sets the build_date of this BuildVersionInfo.


        :param build_date: The build_date of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._build_date = build_date

    @property
    def goarch(self):
        """Gets the goarch of this BuildVersionInfo.  # noqa: E501


        :return: The goarch of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._goarch

    @goarch.setter
    def goarch(self, goarch):
        """Sets the goarch of this BuildVersionInfo.


        :param goarch: The goarch of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._goarch = goarch

    @property
    def goos(self):
        """Gets the goos of this BuildVersionInfo.  # noqa: E501


        :return: The goos of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._goos

    @goos.setter
    def goos(self, goos):
        """Sets the goos of this BuildVersionInfo.


        :param goos: The goos of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._goos = goos

    @property
    def git_commit(self):
        """Gets the git_commit of this BuildVersionInfo.  # noqa: E501


        :return: The git_commit of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._git_commit

    @git_commit.setter
    def git_commit(self, git_commit):
        """Sets the git_commit of this BuildVersionInfo.


        :param git_commit: The git_commit of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._git_commit = git_commit

    @property
    def git_version(self):
        """Gets the git_version of this BuildVersionInfo.  # noqa: E501


        :return: The git_version of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._git_version

    @git_version.setter
    def git_version(self, git_version):
        """Sets the git_version of this BuildVersionInfo.


        :param git_version: The git_version of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._git_version = git_version

    @property
    def major(self):
        """Gets the major of this BuildVersionInfo.  # noqa: E501


        :return: The major of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._major

    @major.setter
    def major(self, major):
        """Sets the major of this BuildVersionInfo.


        :param major: The major of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._major = major

    @property
    def minor(self):
        """Gets the minor of this BuildVersionInfo.  # noqa: E501


        :return: The minor of this BuildVersionInfo.  # noqa: E501
        :rtype: str
        """
        return self._minor

    @minor.setter
    def minor(self, minor):
        """Sets the minor of this BuildVersionInfo.


        :param minor: The minor of this BuildVersionInfo.  # noqa: E501
        :type: str
        """

        self._minor = minor

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(BuildVersionInfo, dict):
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
        if not isinstance(other, BuildVersionInfo):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
