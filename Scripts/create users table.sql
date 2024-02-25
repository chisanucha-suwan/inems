USE [inems_v2];
BEGIN TRANSACTION;
IF NOT EXISTS (
    SELECT 1
    FROM INFORMATION_SCHEMA.TABLES
    WHERE TABLE_NAME = N'Users'
) BEGIN PRINT 'Creating table [inems_v2].[dbo].[Users]...';
CREATE TABLE [dbo].[Users] (
    [ID] [int] IDENTITY(1, 1) NOT NULL PRIMARY KEY,
    [Username] [nvarchar](50) NOT NULL,
    [Password] [nvarchar](max) NOT NULL,
    [PrefixTH] [nvarchar](max) NULL,
    [FirstNameTH] [nvarchar](max) NOT NULL,
    [LastNameTH] [nvarchar](max) NOT NULL,
    [FullNameTH] [nvarchar](max) NOT NULL,
    [PrefixEN] [nvarchar](max) NULL,
    [FirstNameEN] [nvarchar](max) NOT NULL,
    [LastNameEN] [nvarchar](max) NOT NULL,
    [FullNameEN] [nvarchar](max) NOT NULL,
    [Email] [nvarchar](max) NULL,
    [MobilePhone] [nvarchar](max) NULL,
    [ImgProfile] [nvarchar](max) NOT NULL DEFAULT '/images/profiles/default.png',
    [IsMember] [bit] NOT NULL DEFAULT 1,
    [Bio] [nvarchar](max) NULL,
    [BioImageSCV] [nvarchar](max) NULL,
    [IsActive] [bit] NOT NULL DEFAULT 1,
    [Role] [int] NOT NULL DEFAULT 1,
    [PermissionCSV] [nvarchar](max) NULL,
    [CreatedBy] [nvarchar](max) NOT NULL,
    [CreatedDate] [datetime] NOT NULL DEFAULT GETDATE(),
    [ModifiedBy] [nvarchar](max) NULL,
    [ModifiedDate] [datetime] NULL
);
INSERT INTO [dbo].[Users] (
        [Username],
        [Password],
        [PrefixTH],
        [FirstNameTH],
        [LastNameTH],
        [FullNameTH],
        [PrefixEN],
        [FirstNameEN],
        [LastNameEN],
        [FullNameEN],
        [Email],
        [MobilePhone],
        [ImgProfile],
        [IsMember],
        [Bio],
        [BioImageSCV],
        [IsActive],
        [Role],
        [PermissionCSV],
        [CreatedBy],
        [CreatedDate],
        [ModifiedBy],
        [ModifiedDate]
    )
VALUES (
        'administrator',
        '1yhnvnXLIA846DIeg+tflzRLS1BEMmJEYVFBTHBsa0lWNGhiSjFCQWMzTjNNSEpr',
        NULL,
        'ผู้ดูแลระบบ',
        'สูงสุด',
        'ผู้ดูแลระบบ สูงสุด',
        NULL,
        'Super',
        'Administrator',
        'Super Administrator',
        'super_administrator@mail.com',
        NULL,
        '/images/profiles/default.png',
        1,
        NULL,
        NULL,
        1,
        1,
        NULL,
        'System',
        GETDATE(),
        NULL,
        NULL
    );
PRINT 'Table [inems_v2].[dbo].[Users] created.';
END
ELSE BEGIN IF NOT EXISTS (
    SELECT 1
    FROM [dbo].[Users]
    WHERE [Username] = 'administrator'
) BEGIN PRINT 'Username "administrator" does not exist.';
PRINT 'Inserting...';
INSERT INTO [dbo].[Users] (
        [Username],
        [Password],
        [PrefixTH],
        [FirstNameTH],
        [LastNameTH],
        [FullNameTH],
        [PrefixEN],
        [FirstNameEN],
        [LastNameEN],
        [FullNameEN],
        [Email],
        [MobilePhone],
        [ImgProfile],
        [IsMember],
        [Bio],
        [BioImageSCV],
        [IsActive],
        [Role],
        [PermissionCSV],
        [CreatedBy],
        [CreatedDate],
        [ModifiedBy],
        [ModifiedDate]
    )
VALUES (
        'administrator',
        '1yhnvnXLIA846DIeg+tflzRLS1BEMmJEYVFBTHBsa0lWNGhiSjFCQWMzTjNNSEpr',
        NULL,
        'ผู้ดูแลระบบ',
        'สูงสุด',
        'ผู้ดูแลระบบ สูงสุด',
        NULL,
        'Super',
        'Administrator',
        'Super Administrator',
        'super_administrator@mail.com',
        NULL,
        '/images/profiles/default.png',
        1,
        NULL,
        NULL,
        1,
        '1',
        NULL,
        'System',
        GETDATE(),
        NULL,
        NULL
    );
PRINT 'User "administrator" inserted.';
END
ELSE BEGIN PRINT 'Username "administrator" already exists.';
PRINT 'Skipping insert...';
END
END COMMIT TRANSACTION;