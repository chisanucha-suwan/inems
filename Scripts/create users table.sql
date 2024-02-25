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
    [Prefix_TH] [nvarchar](max) NULL,
    [FirstName_TH] [nvarchar](max) NOT NULL,
    [LastName_TH] [nvarchar](max) NOT NULL,
    [FullName_TH] [nvarchar](max) NOT NULL,
    [Prefix_EN] [nvarchar](max) NULL,
    [FirstName_EN] [nvarchar](max) NOT NULL,
    [LastName_EN] [nvarchar](max) NOT NULL,
    [FullName_EN] [nvarchar](max) NOT NULL,
    [Email] [nvarchar](max) NULL,
    [MobliePhone] [nvarchar](max) NULL,
    [ImgProfile] [nvarchar](max) NOT NULL DEFAULT '/images/profiles/default.png',
    [IsMember] [bit] NOT NULL DEFAULT 1,
    [Bio] [nvarchar](max) NULL,
    [BioImageSCV] [nvarchar](max) NULL,
    [IsActive] [bit] NOT NULL DEFAULT 1,
    [Role] [int] NOT NULL DEFAULT 1,
    [PermissionCSV] [nvarchar](max) NULL,
    [Created_By] [nvarchar](max) NOT NULL,
    [Created_Date] [datetime] NOT NULL DEFAULT GETDATE(),
    [Modified_By] [nvarchar](max) NULL,
    [Modified_Date] [datetime] NULL
);
INSERT INTO [dbo].[Users] (
        [Username],
        [Password],
        [Prefix_TH],
        [FirstName_TH],
        [LastName_TH],
        [FullName_TH],
        [Prefix_EN],
        [FirstName_EN],
        [LastName_EN],
        [FullName_EN],
        [Email],
        [MobliePhone],
        [ImgProfile],
        [IsMember],
        [Bio],
        [BioImageSCV],
        [IsActive],
        [Role],
        [PermissionCSV],
        [Created_By],
        [Created_Date],
        [Modified_By],
        [Modified_Date]
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
        [Prefix_TH],
        [FirstName_TH],
        [LastName_TH],
        [FullName_TH],
        [Prefix_EN],
        [FirstName_EN],
        [LastName_EN],
        [FullName_EN],
        [Email],
        [MobliePhone],
        [ImgProfile],
        [IsMember],
        [Bio],
        [BioImageSCV],
        [IsActive],
        [Role],
        [PermissionCSV],
        [Created_By],
        [Created_Date],
        [Modified_By],
        [Modified_Date]
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