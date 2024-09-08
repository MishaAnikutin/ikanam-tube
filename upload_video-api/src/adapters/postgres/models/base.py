from sqlalchemy import MetaData
from sqlalchemy.ext.declarative import as_declarative

metadata = MetaData()


@as_declarative(metadata=metadata)
class BaseTable:
    __allow_unmapped__ = False
